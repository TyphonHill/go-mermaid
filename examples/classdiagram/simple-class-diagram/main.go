package main

import (
	"fmt"

	"github.com/TyphonHill/go-mermaid/classdiagram"
)

func main() {
	cd := classdiagram.NewClassDiagram()
	cd.Title = "Simple Class Diagram"

	cd.AddNote("Test", nil)

	ns := cd.AddNamespace("TestNamespace")

	class1 := cd.AddClass("Class1", ns)
	class1.Label = "ClassLabel"
	class1.Annotation = classdiagram.ClassAnnotationInterface

	cd.AddNote("Class1", class1)

	test := class1.AddMethod("test")
	test.AddParameter("param1", "int")
	test.AddParameter("param2", "string")
	test.ReturnType = "int"
	test.Classifier = classdiagram.MethodClassifierAbstract

	test = class1.AddMethod("test1")
	test.ReturnType = "string"

	test = class1.AddMethod("test2")
	test.Classifier = classdiagram.MethodClassifierStatic
	test.AddParameter("Param1", "List~int~")

	class1.AddField("field1", "string")
	temp := class1.AddField("field2", "string")
	temp.Classifier = classdiagram.FieldClassifierStatic

	class2 := cd.AddClass("Class2", nil)
	class2.Label = "ClassLabel"
	test = class2.AddMethod("test")
	test.AddParameter("param1", "int")
	test.AddParameter("param2", "string")
	test.ReturnType = "int"
	test.Classifier = classdiagram.MethodClassifierAbstract

	test = class2.AddMethod("test1")
	test.ReturnType = "string"

	test = class2.AddMethod("test2")
	test.Classifier = classdiagram.MethodClassifierStatic
	test.AddParameter("Param1", "List~int~")

	class2.AddField("field1", "string")
	temp = class2.AddField("field2", "string")
	temp.Classifier = classdiagram.FieldClassifierStatic

	rel := cd.AddRelation(class1, class2)
	rel.RelationToClassA = classdiagram.RelationTypeAggregation
	rel.RelationToClassB = classdiagram.RelationTypeAssociationLeft
	rel.CardinalityToClassA = classdiagram.RelationCardinalityMany
	rel.CardinalityToClassB = classdiagram.RelationCardinalityMany
	rel.Link = classdiagram.RelationLinkDashed
	rel.Label = "Relation Label"

	fmt.Println(cd.String())
}
