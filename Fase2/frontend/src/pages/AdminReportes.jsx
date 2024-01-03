import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function AdminReportes() {
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo");
        window.open("/principal/admin", "_self");
    };
    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Menu de Reportes</h1>
                    <br />
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" >
                            Reporte Arbol B
                        </button>
                    </center>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" >
                            Reporte Grafo
                        </button>
                    </center>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" >
                            Reporte Arbol Merkle
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
    )
}

export default AdminReportes