import React, { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function Estudiante() {
    const vLibros = (e) => {
        e.preventDefault();
        localStorage.setItem("cursos", JSON.stringify(cursos));
        window.open("/principal/estudiante/ver-libros", "_self");
    };
    const vPublicaciones = (e) => {
        e.preventDefault();
        localStorage.setItem("cursos", JSON.stringify(cursos));
        window.open("/principal/estudiante/ver-publicaciones", "_self");
    };
    const lPrincipal = (e) => {
        e.preventDefault();
        localStorage.clear();
        window.open("/", "_self");
    };
    const [cursos, setCursos] = useState([]);

    useEffect(() => {
        async function PedirCursos() {
            const valorLocal = localStorage.getItem("user");
            const response = await fetch("http://localhost:4000/obtener-clases", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Carnet: valorLocal,
                }),
            });
            const result = await response.json();
            console.log(result);
            setCursos(result.Arreglo);
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

    return (
        <div class="px-4 pt-5 my-5 text-center border-bottom">
            <h1 className="display-4 fw-bold text-body-emphasis bienvenido-titulo">Bienvenido Estudiante</h1>
            <div class="col-lg-6 mx-auto">
                <br />
                <br />
                <p class="lead mb-4">Bienvenido Estudiante, en este menu puedes ver los cursos asignados y las publicaciones de los tutores.
                </p>
                <p class="lead mb-4">Tambien puedes ver los libros subidos al sistema por los tutores de sus respectivos cursos.
                </p>
                <br />
                <br />
                <div class="d-grid gap-2 d-sm-flex justify-content-sm-center mb-5">
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={vLibros}>Ver Libros</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={vPublicaciones}>Ver Publicaciones</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={lPrincipal}>Login Principal</button>

                    <br />
                </div>
                <br />
                <br />
                <h2 className="display-5 fw-bold text-body-emphasis bienvenido-titulo">Cursos Asignados</h2>
                <br />
                {cursos.length > 0 ? <Palabra /> : null}
                <br />
            </div>
        </div>
    )
}

export default Estudiante;
