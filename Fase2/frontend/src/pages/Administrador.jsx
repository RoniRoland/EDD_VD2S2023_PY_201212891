import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function Administrador() {

    const [showSuccessAlertTutor, setShowSuccessAlertTutor] = useState(false);
    const [showSuccessAlertAlumno, setShowSuccessAlertAlumno] = useState(false);
    const [showSuccessAlertCurso, setShowSuccessAlertCurso] = useState(false);

    const reportes = (e) => {
        e.preventDefault();
        window.open("/principal/admin/alumnos", "_self");
    };

    const uploadFileTutor = (event) => {
        const file = event.target.files[0];
        const reader = new FileReader();

        reader.onload = (event) => {
            const content = event.target.result;
            const parsedData = parseCSV(content);
            parsedData.map(async (row) => {
                if (row.length > 1) {
                    const response = await fetch(
                        "http://localhost:4000/registrar-tutor",
                        {
                            method: "POST",
                            headers: {
                                "Content-Type": "application/json",
                            },
                            body: JSON.stringify({
                                Carnet: parseInt(row[0]),
                                Nombre: row[1],
                                Curso: row[2],
                                Password: row[3],
                            }),
                        }
                    );

                    const result = await response.json();
                }
            });
        };

        reader.onerror = (error) => {
            console.error("Error al leer el archivo:", error);
        };

        reader.readAsText(file);
        setShowSuccessAlertTutor(true);
        setTimeout(() => {
            setShowSuccessAlertTutor(false);
        }, 5000);

    };

    const uploadFileAlumno = (event) => {
        const file = event.target.files[0];
        const reader = new FileReader();

        reader.onload = (event) => {
            const content = event.target.result;
            const parsedData = parseCSV(content);
            console.log(parsedData);
            parsedData.map(async (row) => {
                if (row.length > 1) {
                    const response = await fetch(
                        "http://localhost:4000/registrar-alumno",
                        {
                            method: "POST",
                            headers: {
                                "Content-Type": "application/json",
                            },
                            body: JSON.stringify({
                                Carnet: parseInt(row[0]),
                                Nombre: row[1],
                                Password: row[2],
                                Cursos: [row[3], row[4], row[5]],
                            }),
                        }
                    );

                    const result = await response.json();
                }
            });
        };

        reader.onerror = (error) => {
            console.error("Error al leer el archivo:", error);
        };

        reader.readAsText(file);
        setShowSuccessAlertAlumno(true);
        setTimeout(() => {
            setShowSuccessAlertAlumno(false);
        }, 5000);

    };

    const parseCSV = (csvContent) => {
        const rows = csvContent.split("\n");
        const encabezado = rows.slice(1);
        const sinRetorno = encabezado.map((row) => row.trim());
        const data = sinRetorno.map((row) => row.split(","));
        return data;
    };

    const onChange1 = (e) => {
        var reader = new FileReader();
        reader.onload = async (e) => {
            var obj = JSON.parse(e.target.result);
            console.log(obj);
            const response = await fetch("http://localhost:4000/registrar-cursos", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Cursos: obj.Cursos,
                }),
            });
        };
        reader.readAsText(e.target.files[0]);
        setShowSuccessAlertCurso(true);
        setTimeout(() => {
            setShowSuccessAlertCurso(false);
        }, 5000);
    };

    const salir = (e) => {
        e.preventDefault();
        console.log("Listo");
        window.open("/principal/admin", "_self");
    };

    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Menu de Carga</h1>
                    <br />
                    <br />

                    <h2 class="pb-2 border-bottom">Carga de Archivo de Tutores</h2>
                    <div>
                        <br />
                        <input class="form-control form-control-lg" id="formFileLg" type="file"
                            accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
                            onChange={uploadFileTutor} />
                        {showSuccessAlertTutor && (
                            <div className="alert alert-success" role="alert">
                                ¡La carga del archivo de tutor fue exitosa!
                            </div>
                        )}

                    </div>

                    <br />
                    <h2 class="pb-2 border-bottom">Carga de Archivo de Alumnos</h2>
                    <div>
                        <br />
                        <input class="form-control form-control-lg" id="formFileLg" type="file" accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
                            onChange={uploadFileAlumno} />
                        {showSuccessAlertAlumno && (
                            <div className="alert alert-success" role="alert">
                                ¡La carga del archivo de alumno fue exitosa!
                            </div>
                        )}

                    </div>

                    <br />
                    <h2 class="pb-2 border-bottom">Carga de Archivo de Cursos</h2>
                    <div>
                        <br />
                        <input class="form-control form-control-lg" id="formFileLg" type="file" accept="application/json"
                            onChange={onChange1} />
                        {showSuccessAlertCurso && (
                            <div className="alert alert-success" role="alert">
                                ¡La carga del archivo de los Cursos fue exitosa!
                            </div>
                        )}

                    </div>

                    <br />


                    <center>
                        <button className="w-50 btn btn-outline-primary" onClick={reportes}>
                            Ver estudiantes Activos
                        </button>
                    </center>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-success" onClick={salir}>
                            Salir
                        </button>
                    </center>
                    <br />
                </form>
            </div>
        </div>
    );
}

export default Administrador;
