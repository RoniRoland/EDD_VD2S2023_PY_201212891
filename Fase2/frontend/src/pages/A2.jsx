import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'



function A2() {
    const reportes = (e) => {
        e.preventDefault();
        window.open("/principal/admin/reportes", "_self");
    };
    const cArchivos = (e) => {
        e.preventDefault();
        window.open("/principal/admin/cargar-archivos", "_self");
    };
    const cLibro = (e) => {
        e.preventDefault();
        window.open("/principal/admin/aceptar-libros", "_self");
    };
    const lPrincipal = (e) => {
        e.preventDefault();
        window.open("/", "_self");
    };
    return (
        <div class="px-4 pt-5 my-5 text-center border-bottom">
            <h1 className="display-4 fw-bold text-body-emphasis bienvenido-titulo">Bienvenido ADMINISTRADOR</h1>
            <div class="col-lg-6 mx-auto">
                <br />
                <br />
                <p class="lead mb-4">Bienvenido ADMIN 201212891! Para empezar a utilizar el programa puedes elegir cualquiera de las opciones que se encuentra en la parte inferior!
                </p>
                <p class="lead mb-4">Para empezar a utilizar el servicio, elige entre cargar archivos donde se cargaran los archios de los estudiantes, tutores y cursos. Despues puedes elegir la opcion para ver los reportes de los alumnos, cursos y tutores. Por ultimos puedes elegir que libros aceptar o rechazar para los tutores.
                </p>
                <br />
                <br />
                <div class="d-grid gap-2 d-sm-flex justify-content-sm-center mb-5">
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={cArchivos}>Carga de Archivos</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={cLibro}>Libros</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={reportes}>Reportes</button>
                    <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={lPrincipal}>Login Principal</button>

                    <br />
                </div>
                <br />
                <br />
            </div>
        </div>
    )
}

export default A2