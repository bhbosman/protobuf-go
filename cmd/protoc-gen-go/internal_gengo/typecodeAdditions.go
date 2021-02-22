package internal_gengo

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"hash/crc32"
)

var bvisStream = protogen.GoIdent{
	GoName:       "",
	GoImportPath: "github.com/bhbosman/gocommon/stream",
}

var gpProtoExtra = protogen.GoIdent{
	GoName:       "",
	GoImportPath: "github.com/bhbosman/goprotoextra",
}

var bvisConstants = protogen.GoIdent{
	GoName:       "",
	GoImportPath: "github.com/bhbosman/goerrors",
}

var contextPackage = protogen.GoIdent{
	GoName:       "",
	GoImportPath: "context",
}

func genAdditional(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo) {
	for _, message := range f.allMessages {
		if message.Desc.IsMapEntry() {
			continue
		}
		s01 := fmt.Sprintf("%v", message.GoIdent.GoName)
		g.P("// Typecode generated from: ", "\"", s01, "\"")
		g.P("const ", message.GoIdent, "TypeCode uint32 = ", crc32.ChecksumIEEE([]byte(s01)))
		//s02 := fmt.Sprintf(".%v.%v.", message.GoIdent.String(), g.QualifiedGoIdent(message.GoIdent))
		//g.P("// Typecode generated from: ","\"",s02,"\"")
		//g.P("const ", message.GoIdent, "TypeCode uint32 = ", crc32.ChecksumIEEE([]byte(s02)))
	}

	for _, message := range f.allMessages {
		if message.Desc.IsMapEntry() {
			continue
		}
		g.P("//", message.genRawDescMethod)
		g.P("//", message.genExtRangeMethod)
		g.P("//", message.isTracked)
		g.P("//", message.hasWeak)

		g.P("type ", message.GoIdent, "Wrapper struct {")
		g.P(g.QualifiedGoIdent(gpProtoExtra), "BaseMessageWrapper")
		g.P("Data *", message.GoIdent)
		g.P("}")
		g.P()
		g.P("func (self *", message.GoIdent, "Wrapper) Message() interface{} {")
		g.P("return self.Data")
		g.P("}")
		g.P()
		g.P("func (self *", message.GoIdent, "Wrapper) messageWrapper()", " interface{} {")
		g.P("return self")
		g.P("}")
		g.P()
		g.P("func New", message.GoIdent, "Wrapper(")
		g.P("cancelCtx ", g.QualifiedGoIdent(contextPackage), "Context,")
		g.P("cancelFunc ", g.QualifiedGoIdent(contextPackage), "CancelFunc,")
		g.P("toReactor ", g.QualifiedGoIdent(gpProtoExtra), "ToReactorFunc,")
		g.P("toConnection ", g.QualifiedGoIdent(gpProtoExtra), "ToConnectionFunc,")
		g.P("data *", message.GoIdent, ") *", message.GoIdent, "Wrapper {")
		g.P("return &", message.GoIdent, "Wrapper{")
		g.P("BaseMessageWrapper: ", g.QualifiedGoIdent(gpProtoExtra), "NewBaseMessageWrapper(")
		g.P("cancelCtx,")
		g.P("cancelFunc,")
		g.P("toReactor,")
		g.P("toConnection),")
		g.P("Data: data,")
		g.P("}")
		g.P("}")
		g.P("")

		g.P("var _ = ", g.QualifiedGoIdent(bvisStream), "Register(")
		g.P(message.GoIdent, "TypeCode,")
		g.P(g.QualifiedGoIdent(bvisStream), "TypeCodeData{")
		g.P("CreateMessage: func() proto1.Message {")
		g.P("return &", message.GoIdent, "{}")
		g.P("},")
		g.P("CreateWrapper: func(")
		g.P("cancelCtx ", g.QualifiedGoIdent(contextPackage), "Context,")
		g.P("cancelFunc ", g.QualifiedGoIdent(contextPackage), "CancelFunc,")
		g.P("toReactor ", g.QualifiedGoIdent(gpProtoExtra), "ToReactorFunc,")
		g.P("toConnection ", g.QualifiedGoIdent(gpProtoExtra), "ToConnectionFunc,")
		g.P("data ", protoPackage.Ident("Message"), ") (", g.QualifiedGoIdent(gpProtoExtra), "IMessageWrapper, error) {")
		g.P("if msg, ok := data.(*", message.GoIdent, "); ok {")
		g.P("return New", message.GoIdent, "Wrapper(")
		g.P("cancelCtx,")
		g.P("cancelFunc,")
		g.P("toReactor,")
		g.P("toConnection,")
		g.P("msg), nil")
		g.P("}")
		g.P("return nil, ", g.QualifiedGoIdent(bvisConstants), "InvalidParam")
		g.P("}})")

	}
}

func genMessageTypeCodeMethod(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	g.P("func (self *", m.GoIdent, ") TypeCode() uint32 {")
	g.P("return ", m.GoIdent, "TypeCode")
	g.P("}")
	g.P()
}
