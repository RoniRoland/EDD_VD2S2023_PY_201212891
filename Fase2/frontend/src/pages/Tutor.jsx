import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function Tutor() {
  const cLibro = (e) => {
    e.preventDefault();
    window.open("/principal/tutor/libro", "_self");
  };
  const cPublicacion = (e) => {
    e.preventDefault();
    window.open("/principal/tutor/publicacion", "_self");
  };
  const lPrincipal = (e) => {
    e.preventDefault();
    window.open("/", "_self");
  };
  return (
    <div class="px-4 pt-5 my-5 text-center border-bottom">
      <h1 className="display-4 fw-bold text-body-emphasis bienvenido-titulo">Bienvenido Tutor</h1>
      <div class="col-lg-6 mx-auto">
        <br />
        <br />
        <p class="lead mb-4">Bienvenido Tutor, en este menu puedes cargar un libro en formato pdf y ver si el estado esta en aceptado o rechazado por el administrador.
        </p>
        <p class="lead mb-4">Tambien puedes realizar publicaciones en el apartado de cargar publicacion.
        </p>
        <br />
        <br />
        <div class="d-grid gap-2 d-sm-flex justify-content-sm-center mb-5">
          <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={cLibro}>Carga de Libros</button>
          <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={cPublicacion}>Carga de Publicacion</button>
          <button type="button" class="btn btn-primary btn-lg px-4 me-sm-3" onClick={lPrincipal}>Login Principal</button>

          <br />
        </div>
        <br />
        <br />
      </div>

    </div>
  )
}

export default Tutor