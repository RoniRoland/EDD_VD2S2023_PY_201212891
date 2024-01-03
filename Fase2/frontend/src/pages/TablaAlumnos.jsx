import React, { useState, useEffect } from "react";

function TablaAlumnos() {
  const [alumnosRegistrados, SetAlumnosRegistrados] = useState([]);
  useEffect(() => {
    async function peticion() {
      const response = await fetch("http://localhost:4000/tabla-alumnos");
      const result = await response.json();
      SetAlumnosRegistrados(result.Arreglo);
    }
    peticion();
  }, []);

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/admin/cargar-archivos", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Estudiantes Activos</h1>

          <br />
          <br />
          <table className="table table-dark table-bordered border-info-subtle table-striped">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Posicion</th>
                <th scope="col">Carnet </th>
                <th scope="col">Password </th>
                <th scope="col">Curso 1 </th>
                <th scope="col">Curso 2 </th>
                <th scope="col">Curso 3 </th>
              </tr>
            </thead>
            <tbody>
              {alumnosRegistrados.map((element, j) => {
                if (element.Id_Cliente != "") {
                  return (
                    <>
                      <tr key={"alum" + j}>
                        <th scope="row">{j + 1}</th>
                        <td>{element.Llave}</td>
                        <td>{element.Persona.Carnet}</td>
                        <td>{element.Persona.Password}</td>
                        <td>{element.Persona.Cursos[0]}</td>
                        <td>{element.Persona.Cursos[1]}</td>
                        <td>{element.Persona.Cursos[2]}</td>
                      </tr>
                    </>
                  );
                }
              })}
            </tbody>
          </table>
          <br />
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <p className="mt-5 mb-3 text-muted">EDD 201212891</p>
          <br />
        </form>
      </div>

    </div>
  );
}

export default TablaAlumnos;
