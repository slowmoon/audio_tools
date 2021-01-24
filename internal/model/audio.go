package model

import "gorm.io/gorm"

type Audios struct {
    gorm.Model
    Avatar  string
    Name    string
    Type    string
}

func (a Audios)TableName() string {
    return  "audios_table"
}
