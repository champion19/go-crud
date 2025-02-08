package main

import (
    "bytes"
    "io"
    "os"
    "testing"
)

func TestMainOutput(t *testing.T) {
    // Guardamos la salida estándar original
    originalStdout := os.Stdout
    defer func() { os.Stdout = originalStdout }() // Restaurar al final

    // Creamos un buffer para capturar la salida
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Ejecutamos main()
    main()

    // Cerramos el writer y leemos la salida
    w.Close()
    var buf bytes.Buffer
    io.Copy(&buf, r)

    // Comprobamos el resultado
    esperado := "Hi, World!\n"
    resultado := buf.String()
    if resultado != esperado {
        t.Errorf("Esperado: %q, Obtenido: %q", esperado, resultado)
    }
}
