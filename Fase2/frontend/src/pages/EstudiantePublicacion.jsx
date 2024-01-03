import React, { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function EstudiantePublicacion() {
    const [cursos, setCursos] = useState([]);
    const [publicacion, setPublicacion] = useState([]);

    useEffect(() => {
        async function PedirCursos() {
            const arregloCursos = localStorage.getItem("cursos");
            const cursosArreglo = JSON.parse(arregloCursos);
            const response = await fetch(
                "http://localhost:4000/obtener-publi-alumno"
            );
            const result = await response.json();
            console.log(result);
            setCursos(cursosArreglo);
            if (result.status === 200) {
                setPublicacion(result.Arreglo);
            }
        }
        PedirCursos();
    }, []);

    const Palabra = () => {
        return (
            <div className="row">
                <div className="row align-items-start">
                    {cursos.map((item, i) => (
                        <div className="form-signin1 col" key={"CursoEstudiante" + i}>
                            <div className="text-center">
                                <div className="card card-body">
                                    <h1 className="text-left" key={"album" + i} value={i}>
                                        {item}
                                    </h1>
                                    <div>
                                        <span
                                            className="input-group-text"
                                            id="validationTooltipUsernamePrepend"
                                        ></span>{" "}
                                        <br />
                                        {publicacion.map((ite, j) => {
                                            if (ite.Curso === item) {
                                                return (
                                                    <div key={"p-" + i}>
                                                        <label className="input-group-text" key={"lib" + j}>
                                                            {ite.Contenido}
                                                        </label>
                                                    </div>
                                                );
                                            }
                                        })}
                                    </div>
                                </div>
                                <br />
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        );
    };

    const salir = (e) => {
        e.preventDefault();
        window.open("/principal/estudiante", "_self");
    };

    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Publicaciones de Tutores en los Cursos</h1>
                    <br />
                    <br />
                    {cursos.length > 0 ? <Palabra /> : null}
                    <br />

                    <center>
                        <button className="w-50 btn btn-outline-primary" onClick={salir}>
                            Salir
                        </button>
                    </center>
                </form>
            </div>
        </div>
    )
}

export default EstudiantePublicacion