package Utilidades

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

//FUnción CrearArchivo
func CrearArchivo(name string) error {
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error en la dirección del Archivo: ", err)
		return err
	}

	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Error en la creación de un Archivo: ", err)
			return err
		}
		defer file.Close()
	}
	return nil
}

//Funcion Abrir Archivo
func AbrirArchivo(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return nil, err
	}
	return file, nil
}

// Funcion EcribirObjeto
func EscribirObjeto(file *os.File, data interface{}, position int64) error {
	file.Seek(position, 0)
	err := binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Error al escribir un objeto en el archivo: ", err)
		return err
	}
	return nil
}

// Funcion LeerObjeto
func LeerObjeto(file *os.File, data interface{}, position int64) error {
	file.Seek(position, 0)
	err := binary.Read(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Error al leer el objeto en el archivo: ", err)
		return err
	}
	return nil
}