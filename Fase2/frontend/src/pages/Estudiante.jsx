import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function Estudiante() {
    const vCursos = (e) => {
        e.preventDefault();
        window.open("/principal/estudiante/ver-cursos", "_self");
    };
    const vLibros = (e) => {
        e.preventDefault();
        window.open("/principal/estudiante/ver-libros", "_self");
    };
    const vPublicaciones = (e) => {
        e.preventDefault();
        window.open("/principal/estudiante/ver-publicaciones", "_self");
    };
    const lPrincipal = (e) => {
        e.preventDefault();
        window.open("/", "_self");
    };
    return (
        <div class="px-4 pt-5 my-5 text-center border-bottom">
            <h1 className="display-4 fw-bold text-body-emphasis bienvenido-titulo">Bienvenido Estudiante</h1>
            <div class="col-lg-6 mx-auto">
                <br />
                <br />
                <p class="lead mb-4">Bienvenido Estudiante, en este menu puedes ver los cursos asignados y las publicaciones de los tutores.
                </p>
                <p class="lead mb-4">Tambien puedes ver y descargar los libros subidos al sistema por los tutores.
                </p>
                <br />
                <br />
                <div class="d-grid gap-2 d-sm-flex justify-content-sm-center mb-5">
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={vCursos}>Ver Cursos</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={vLibros}>Ver Libros</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={vPublicaciones}>Ver Publicaciones</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={lPrincipal}>Login Principal</button>

                    <br />
                </div>
                <br />
                <br />
            </div>
        </div>
    )
}

export default Estudiante;
