package ast

import (
	"strings"
	"unicode"

	"github.com/lzy0505/gocap/generator"

	"github.com/lzy0505/gocap/utils"
)

type Typ interface {
	ToString() string
}

type IntType struct{}

func (t IntType) ToString() string {
	return "int"
}

type StringType struct{}

func (t StringType) ToString() string {
	return "string"
}

type NamedType struct {
	TypeId string
}

func (t NamedType) ToString() string {
	return t.TypeId
}

type ImportedType struct {
	PackageId string
	typeId    string
}

func (t ImportedType) ToString() string {
	return t.PackageId + "." + t.typeId
}

type StructType struct {
	Fields []StructField
}

func (t StructType) ToString() string {
	result := "struct { \n"
	for _, f := range t.Fields {
		result += f.ToString()
	}
	return result + "}\n"
}

type StructField struct {
	Id  string
	Typ Typ
}

func (t StructField) ToString() string {
	return t.Id + " " + t.Typ.ToString() + "\n"
}

type PointerType struct {
	Typ Typ
}

func (t PointerType) ToString() string {
	return "*" + t.Typ.ToString()
}

type FunctionType struct {
	Params     []Typ
	ReturnType []Typ
}

func (t FunctionType) ToString() string {
	s := "func("
	if len(t.Params) > 0 {
		s += t.Params[0].ToString()
		for _, typ := range t.Params[1:] {
			s += ", " + typ.ToString()
		}
	}
	s += ") "

	if len(t.ReturnType) == 1 {
		s += t.ReturnType[0].ToString()
	} else if len(t.ReturnType) > 1 {
		s += "(" + t.Params[0].ToString()
		for _, typ := range t.Params[1:] {
			s += ", " + typ.ToString()
		}
		s += ")"
	}
	return s
}

type InterfaceType struct {
	methods []InterfaceMethod
}

func (t InterfaceType) ToString() string {
	result := "interface { \n"
	for _, m := range t.methods {
		result += m.ToString() + "\n"
	}
	return result + "}\n"
}

type InterfaceMethod struct {
	id        string
	signature Signature
}

func (t InterfaceMethod) ToString() string {
	return t.id + " " + t.signature.ToString()
}

type SliceType struct {
	Typ Typ
}

func (t SliceType) ToString() string {
	return "[]" + t.Typ.ToString()
}

type MapType struct {
	key  Typ
	Elem Typ
}

func (t MapType) ToString() string {
	return "map[" + t.key.ToString() + "]" + t.Elem.ToString()
}

type ChannelType struct {
	Typ Typ
}

func (t ChannelType) ToString() string {
	return "chan " + t.Typ.ToString()
}

type ROChannelType struct {
	typ Typ
}

func (t ROChannelType) ToString() string {
	return "<-chan " + t.typ.ToString()
}

type SOChannelType struct {
	typ Typ
}

func (t SOChannelType) ToString() string {
	return "chan<- " + t.typ.ToString()
}

type CapChannelType struct {
	Typ Typ
}

var typeCapChannelTemplate = "capchan.Type_$TYPE"

func (t CapChannelType) ToString() string {
	switch captyp := t.Typ.(type) {
	case ImportedType:
		typename, ok := generator.ExportedTypeMap[captyp.PackageId]
		if ok {
			return captyp.PackageId + "." + "Type_" + typename
		}
	case PointerType:
		switch captyp2 := captyp.Typ.(type) {
		case ImportedType:
			typename, ok := generator.ExportedTypeMap[captyp2.PackageId]
			if ok {
				return captyp2.PackageId + "." + "Type__st_" + typename
			}
		case NamedType:
			if unicode.IsUpper(rune(captyp2.TypeId[0])) {
				return "Type__st_" + captyp2.TypeId
			}
		}
	case NamedType:
		if unicode.IsUpper(rune(captyp.TypeId[0])) {
			return "Type_" + captyp.TypeId
		}
	}
	//generator.ImportPackage = append(generator.ImportPackage, utils.TempPkg)

	typeString := utils.RemoveParentheses(t.Typ.ToString())
	result := strings.Replace(typeCapChannelTemplate, "$TYPE", typeString, -1)
	return result
}

//TODO: doesn't support RO SO capchannels
type ROCapChannelType struct {
	typ Typ
}
type SOCapChannelType struct {
	typ Typ
}

func NewNamedType(id_ Attrib) (NamedType, error) {
	id := parseId(id_)
	return NamedType{id}, nil
}

func NewStructType(fields_ Attrib) (StructType, error) {
	fields := fields_.([]StructField)
	return StructType{fields}, nil
}

func MakeStructFields(idlist_, typ_ Attrib) ([]StructField, error) {
	idlist := idlist_.([]string)
	typ := typ_.(Typ)
	fields := make([]StructField, len(idlist))
	for i, id := range idlist {
		fields[i] = StructField{id, typ}
	}
	return fields, nil
}

func NewStructFieldList(list_ Attrib) ([]StructField, error) {
	list := list_.([]StructField)
	return list, nil
}

func AppendStructFields(fielddecls1_, fielddecls2_ Attrib) ([]StructField, error) {
	fielddecls1 := fielddecls1_.([]StructField)
	fielddecls2 := fielddecls2_.([]StructField)
	return append(fielddecls1, fielddecls2...), nil
}

func NewTypeList(typ_ Attrib) ([]Typ, error) {
	typ := typ_.(Typ)
	list := make([]Typ, 1)
	list[0] = typ
	return list, nil
}

func AppendTypeList(typlist_, typ_ Attrib) ([]Typ, error) {
	typlist := typlist_.([]Typ)
	typ := typ_.(Typ)
	return append(typlist, typ), nil
}

func NewPointerType(baseType_ Attrib) (PointerType, error) {
	baseType := baseType_.(Typ)
	return PointerType{baseType}, nil
}

func NewFunctionType(params_, result_ Attrib) (FunctionType, error) {
	params := params_.([]Typ)
	result := result_.([]Typ)
	return FunctionType{params, result}, nil
}

func NewImportedType(id1_, id2_ Attrib) (ImportedType, error) {
	id1 := parseId(id1_)
	id2 := parseId(id2_)
	return ImportedType{id1, id2}, nil
}

func NewInterfaceType(methods_ Attrib) (InterfaceType, error) {
	methods := methods_.([]InterfaceMethod)
	return InterfaceType{methods}, nil
}

func NewInterfaceMethod(id_, signature_ Attrib) (InterfaceMethod, error) {
	id := parseId(id_)
	signature := signature_.(Signature)
	return InterfaceMethod{id, signature}, nil
}

func NewInterfaceMethodList(method_ Attrib) ([]InterfaceMethod, error) {
	method := method_.(InterfaceMethod)
	list := make([]InterfaceMethod, 1)
	list[0] = method
	return list, nil
}

func AppendInterfaceMethodList(list_, method_ Attrib) ([]InterfaceMethod, error) {
	list := list_.([]InterfaceMethod)
	method := method_.(InterfaceMethod)
	return append(list, method), nil
}

func NewSliceType(typ_ Attrib) (SliceType, error) {
	typ := typ_.(Typ)
	return SliceType{typ}, nil
}

func NewMapType(keytyp_, elemtyp_ Attrib) (MapType, error) {
	key := keytyp_.(Typ)
	elem := elemtyp_.(Typ)
	return MapType{key, elem}, nil
}

func NewChannelType(typ_ Attrib) (ChannelType, error) {
	typ := typ_.(Typ)
	return ChannelType{typ}, nil
}

func NewROChannelType(typ_ Attrib) (ROChannelType, error) {
	typ := typ_.(Typ)
	return ROChannelType{typ}, nil
}

func NewSOChannelType(typ_ Attrib) (SOChannelType, error) {
	typ := typ_.(Typ)
	return SOChannelType{typ}, nil
}

func NewCapChanType(typ_ Attrib) (CapChannelType, error) {
	typ := typ_.(Typ)
	return CapChannelType{typ}, nil
}

func NewROCapChanType(typ_ Attrib) (ROCapChannelType, error) {
	typ := typ_.(Typ)
	return ROCapChannelType{typ}, nil
}

func NewSOCapChanType(typ_ Attrib) (SOCapChannelType, error) {
	typ := typ_.(Typ)
	return SOCapChannelType{typ}, nil
}

// Declarations
type TypeDeclBlock struct {
	Decls []Code
}

func (t TypeDeclBlock) ToString() string {
	s := t.Decls[0].ToString() + "\n"
	for _, decl := range t.Decls[1:] {
		s += decl.ToString() + "\n"
	}
	return s
}

type TypeDecl struct {
	Id  string
	Typ Typ
}

func (t TypeDecl) ToString() string {
	return "type " + t.Id + " " + t.Typ.ToString()
}

type TypeAlias struct {
	Id  string
	Typ Typ
}

func (t TypeAlias) ToString() string {
	return "type " + t.Id + " = " + t.Typ.ToString()
}

func NewTypeDeclBlock(decls_ Attrib) (TypeDeclBlock, error) {
	decls := decls_.([]Code)
	return TypeDeclBlock{decls}, nil
}

func NewTypeDecl(id_, typ_ Attrib) (TypeDecl, error) {
	id := parseId(id_)
	typ := typ_.(Typ)
	return TypeDecl{id, typ}, nil
}

func NewTypeAlias(id_, typ_ Attrib) (TypeAlias, error) {
	id := parseId(id_)
	typ := typ_.(Typ)
	return TypeAlias{id, typ}, nil
}

func NewTypeSpecList(typespec_ Attrib) ([]Code, error) {
	typespec := typespec_.(Code)
	list := make([]Code, 1)
	list[0] = typespec
	return list, nil
}

func AppendTypeSpecs(types1_, types2_ Attrib) ([]Code, error) {
	types1 := types1_.([]Code)
	types2 := types2_.(Code)
	return append(types1, types2), nil
}
