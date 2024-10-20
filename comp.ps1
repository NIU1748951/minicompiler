# Directorios del proyecto
$binDir = "bin"       # Carpeta donde se guardarán los binarios
$srcDir = "src"       # Carpeta del código fuente
$outDir = "out"       # Carpeta de salida donde se generan otros archivos (ej. ensamblador)
$exeName = "koko.exe" # Nombre del archivo ejecutable

# Asegurar que las carpetas existan
if (-Not (Test-Path $binDir)) {
    New-Item -ItemType Directory -Path $binDir
}
if (-Not (Test-Path $outDir)) {
    New-Item -ItemType Directory -Path $outDir
}

# Ruta del archivo ejecutable
$exePath = Join-Path $binDir $exeName

# Función para limpiar archivos generados
function Clean {
    Write-Host "Limpiando archivos generados..."
    Remove-Item $exePath -Force -ErrorAction SilentlyContinue
    Remove-Item "$outDir\*" -Force -ErrorAction SilentlyContinue
    Write-Host "Limpieza completa."
}

# Función de limpieza si se proporciona el argumento
if ($args[0] -eq "clean") {
    Clean
    exit
}

# Verificar si el archivo ejecutable ya existe y borrarlo
if (Test-Path $exePath) {
    Write-Host "Eliminando ejecutable anterior: $exePath"
    Remove-Item $exePath
} else {
    Write-Host "No se encontro ejecutable previo: $exePath"
}

# Compilar el programa Go y generar el ejecutable en la carpeta bin
Write-Host "Compilando el archivo Go..."
go build -o $exePath ".\$srcDir\."

# Verificar si la compilación fue exitosa
if (Test-Path $exePath) {
    Write-Host "Compilacion exitosa. Se ha creado el archivo $exePath"
} else {
    Write-Host "Error en la compilacion."
}

if ($args[0] -eq "run") {
    # Almacenar el nombre del archivo de entrada
    $inputFile = $args[1]

    # Ejecutar el archivo .exe con el archivo de entrada como argumento
    Write-Host "Ejecutando $exePath con el archivo de entrada '$inputFile'..."
    Write-Host ""
    & $exePath $inputFile
    Write-Host ""
    # Verificar el resultado de la ejecución
    if ($LASTEXITCODE -ne 0) {
        Write-Host "El programa termino con errores."
    } else {
        Write-Host "El programa se ejecutó correctamente."
    }
}