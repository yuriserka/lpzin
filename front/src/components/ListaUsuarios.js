import React, { Component } from 'react'
import PropTypes from 'prop-types'
import ImagemPerfil from './ImagemPerfil';

export class ListaUsuarios extends Component {
  render() {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', flexWrap: 'wrap', overflowY: 'scroll' }}>
        {
          this.props.usuariosAtivos.map((ua) => (
            ua.ID !== this.props.myID &&
            <div key={ua.ID} style={{
              background: 'white', borderRadius: '10px', height: '50px',
              margin: '5px', width: '50%'
            }}>
              <div style={{ display: 'flex' }}>
                <ImagemPerfil obj={ua} h={30} w={30} />
                <span style={{ padding: '10px 0 0 10px' }}>
                  {ua.Nome}
                </span>
                {/* <div style={{ float: 'right' }}>
                  <input type="checkbox" style={{ justifyContent: 'flex-end' }} />
                </div> */}
              </div>
            </div>
          ))
        }
      </div>
    )
  }
}

ListaUsuarios.propTypes = {
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
}

export default ListaUsuarios
