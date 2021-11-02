#!/usr/bin/env python
# -*- coding: utf-8 -*-
import clang.cindex
import os
# import json

StructFile = './include/ThostFtdcUserApiStruct.h'
DataTypeFile = './include/ThostFtdcUserApiDataType.h'
Structs = []
typeMap = {}


class Struct(dict):
    def __init__(self, name, **fields):
        "docstring"
        self.__dict__ = self
        self.name = name
        self.fields = {}
        for k, v in fields.items():
            self.fields[k] = v

    def upper(self, word):
        return word[:1].upper() + word[1:]

    def generate_c_header(self):
        buf = 'typedef struct %s %s;\n' % (self.name, self.name)
        buf += '%s* new%s();\n\n' % (self.name, self.name)
        return buf

    def generate_c_file(self):
        buf = '%s* new%s(){\n' % (self.name, self.name)
        buf += '    %s *ptr = (%s*)malloc(sizeof(%s));\n' % (self.name, self.name, self.name)
        buf += '    return ptr;\n}\n'
        return buf

    def generate(self):
        buf = self.generate_struct()
        buf += '\n'
        buf += self.generate_new()
        return buf

    def generate_struct(self):
        buf = 'type %s struct{\n' % (self.upper(self.name))
        for k, v in self.fields.items():
            buf += '%s %s\n' % (self.upper(k), self.get_type(v))
        buf += '}'
        return buf

    def generate_new(self):
        gname = self.upper(self.name)
        buf = 'func New%s(p *C.%s) *%s{\n' % (gname, self.name, gname)
        buf += 'ret := new(%s)\n' % gname
        for k, v in self.fields.items():
            typ = self.get_type(v)
            if typ == 'int':
                buf += 'ret.%s = int(p.%s)\n' % (self.upper(k), k)
            elif typ == 'int16':
                buf += 'ret.%s = int16(p.%s)\n' % (self.upper(k), k)
            elif typ == 'byte':
                buf += 'ret.%s = byte(p.%s)\n' % (self.upper(k), k)
            elif typ == 'string':
                buf += 'ret.%s = c2goStr(&p.%s[0], C.sizeof_%s)\n' % (self.upper(k), k, v)
            elif typ == 'float64':
                buf += 'ret.%s = goFloat64(p.%s)\n' % (self.upper(k), k)
        buf += 'return ret'
        buf += '}\n'
        buf += self.generate_cvalue()
        return buf

    def generate_cvalue(self):
        gname = self.upper(self.name)
        buf = 'func (s *%s)CValue()* C.%s{\n' % (gname, self.name)
        buf += 'ptr := C.new%s()\n' % self.name
        for k, v in self.fields.items():
            typ = self.get_ctype(v)
            if typ == 'int':
                buf += 'ptr.%s = C.int(s.%s)\n' % (k, self.upper(k))
            elif typ == 'short':
                buf += 'ptr.%s = C.short(s.%s)\n' % (k, self.upper(k))
            elif typ == 'char':
                buf += 'ptr.%s = C.char(s.%s)\n' % (k, self.upper(k))
            elif typ == 'char_array':
                buf += 'go2cStr(s.%s, &ptr.%s[0], C.sizeof_%s)\n' % (self.upper(k), k, v)
            elif typ == 'double':
                buf += 'ptr.%s = C.double(s.%s)\n' % (k, self.upper(k))
        buf += 'return ptr'
        buf += '}\n'
        return buf

    def get_ctype(self, typ):
        ctype = typeMap[typ]
        if 'char [' in ctype:
            return 'char_array'
        return ctype

    def get_type(self, typ):
        ctype = typeMap[typ]
        if ctype == 'int':
            return 'int'
        elif ctype == 'short':
            return 'int16'
        elif ctype == 'char':
            return 'byte'
        elif 'char [' in ctype:
            return 'string'
        elif ctype == 'double':
            return 'float64'
        else:
            return ctype


def analyzer_struct(node, f):
    if not node.location.file or node.location.file.name != f:
        return
    text = node.spelling or node.displayname
    kind = str(node.kind)[str(node.kind).index('.')+1:]
    if kind != "STRUCT_DECL":
        return
    fields = {}
    for c in node.get_children():
        field_type = c.type.spelling
        fields[c.spelling] = field_type
    Structs.append(Struct(text, **fields))


def analyzer_typedef(node, f):
    if not node.location.file or node.location.file.name != f:
        return
    text = node.spelling or node.displayname
    kind = str(node.kind)[str(node.kind).index('.')+1:]
    if kind != "TYPEDEF_DECL":
        return
    typeMap[text] = node.underlying_typedef_type.spelling


def walk(node, fn, f):
    fn(node, f)
    for c in node.get_children():
        walk(c, fn, f)


def generate_c():
    header = open("types_gen.h", "w")
    source = open("types_gen.c", "w")
    header.write('#ifndef _CTP_TYPES_GEN_H\n')
    header.write('#define _CTP_TYPES_GEN_H\n')
    header.write('#include "include/ThostFtdcUserApiStruct.h"\n')
    source.write('#include <stdlib.h>\n')
    source.write('#include "types_gen.h"\n')
    for v in Structs:
        header.write(v.generate_c_header())
        source.write(v.generate_c_file())
    header.write('#endif /* _CTP_TYPES_GEN_H */\n')
    header.close()
    source.close()

def generate_go():
    output = "types_gen.go"
    f = open(output, "w")
    f.write('package ctp\n')
    f.write('/*\n')
    f.write('#include "types_gen.h"\n')
    f.write('\n*/\nimport "C"\n')
    for v in Structs:
        buf = v.generate()
        f.write(buf)
    f.close()
    os.system("gofmt -w " + output)

def main():
    idx = clang.cindex.Index.create()
    tu = idx.parse(DataTypeFile, args=['-std=c++11'], options=0)
    walk(tu.cursor, analyzer_typedef, DataTypeFile)

    tu = idx.parse(StructFile, args=['-std=c++11'], options=0)
    walk(tu.cursor, analyzer_struct, StructFile)
    # for i, v in typeMap.items():
    #   print(i, v)

    generate_c()
    generate_go()


if __name__ == '__main__':
    main()
