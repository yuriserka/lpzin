import React, { Component } from 'react'
import PropTypes from 'prop-types'
import ImagemPerfil from './ImagemPerfil';

export class ListaUsuarios extends Component {
  selecionarUsuario(u) {
    u.Selecionado = !u.Selecionado
    this.props.selecionar(u)
  }

  getUsuarioDivStyle = () => {
    return {
      background: 'white',
      borderRadius: '10px',
      height: '50px',
      margin: '5px', width: '50%'
    }
  }

  render() {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', flexWrap: 'wrap' }}>
        {
          this.props.usuariosAtivos.map(u => { u.Selecionado = false; return u }).map(ua => (
            ua.ID !== this.props.myID &&
            <div key={ua.ID} style={this.getUsuarioDivStyle()}>
              <div style={{ display: 'flex' }}>
                <ImagemPerfil obj={ua} h={30} w={30} />
                <span style={{ padding: '10px 0 0 10px' }}>
                  {ua.Nome}
                </span>
                <div>
                  <input type="checkbox" onChange={this.selecionarUsuario.bind(this, ua)} />
                </div>
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
  selecionar: PropTypes.func.isRequired,
  myID: PropTypes.number.isRequired,
}

export default ListaUsuarios
