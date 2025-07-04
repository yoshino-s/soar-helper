package main

import (
	"bytes"
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/yoshino-s/entproto"
	"github.com/yoshino-s/go-framework/utils"
)

func main() {
	ext, _ := entproto.NewExtension(
		entproto.SkipGenFile(),
		entproto.WithProtoDir("./proto"),
	)
	if err := entc.Generate("./internal/ent/schema/",
		&gen.Config{
			Features: []gen.Feature{
				gen.FeatureExecQuery,
				gen.FeatureUpsert,
				gen.FeatureModifier,
				gen.FeatureIntercept,
				gen.FeatureVersionedMigration,
			},
			Templates: []*gen.Template{
				utils.Must(
					gen.NewTemplate("extend").ParseFiles(
						"./internal/ent/extend.tmpl",
					),
				),
			},
		},
		entc.Extensions(
			ext,
		),
	); err != nil {
		log.Fatal("running ent code gen:", err)
	}

	dir, _ := os.ReadDir("./proto/entpb")
	for _, d := range dir {
		content, _ := os.ReadFile("./proto/entpb/" + d.Name())
		content = bytes.ReplaceAll(content, []byte("ent/proto/entpb"), []byte("proto/entpb"))

		if err := os.WriteFile("./proto/entpb/"+d.Name(), content, 0644); err != nil {
			log.Fatal("writing file:", err)
		}
	}
}
