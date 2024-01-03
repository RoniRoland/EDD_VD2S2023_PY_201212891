import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function PublicacionTutor() {
    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Carga de Publicaciones</h1>
                    <br />
                    <br />
                    <div class="form-floating">
                        <textarea class="form-control" placeholder="Leave a comment here" id="floatingTextarea" ></textarea>
                        <label for="floatingTextarea">Comments</label>
                    </div>
                    <br />
                    <br />
                    <button type="button" class="btn btn-success">Cargar Publicacion</button>
                </form>
            </div>
        </div>
    )
}

export default PublicacionTutor