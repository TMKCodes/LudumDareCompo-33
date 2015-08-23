package main

import (
	"log"
	"io/ioutil"
	sf "bitbucket.org/krepa098/gosfml2"
)


type Texture struct {
	Name string
	Data *sf.Texture
}

func NewTextures(dirname string) []Texture {
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	Textures := make([]Texture, 0)
	for _, FileInfo := range(dir) {
		var Tex Texture
		Tex.Name = FileInfo.Name()
		Tex.Data, err = sf.NewTextureFromFile(dirname + FileInfo.Name(), nil)
		if err != nil {
			log.Fatal(err)
		}
		Textures = append(Textures, Tex)
	}
	return Textures
}

func GetTexture(TextureName string, Textures []Texture) Texture {
	for Texture, _ := range Textures {
		if Textures[Texture].Name == TextureName {
			return Textures[Texture]
		}
	}
	return Textures[0]
}