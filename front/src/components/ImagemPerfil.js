import React from 'react';
import PropTypes from 'prop-types';

export class ImagemPerfil extends React.Component {
  ContainerImagem = () => {
    return {
      display: 'flex',
      width: '55px',
      height: '55px',
      overflow: 'hidden',
      borderRadius: '50%',
      zIndex: '1',
    }
  }

  getImageStyle = () => {
    return {
      width: 'auto',
      height: '100%',
    }
  }

  getFotoPerfil = () => {
    const temFotoPerfil = this.props.obj.FotoPerfil === '' ? false : true
    if (!temFotoPerfil) {
      if (this.props.ehGrupo) {
        return <img src={process.env.PUBLIC_URL + '/images/defaultGrupo.jpg'} alt="" style={this.getImageStyle()} />
      }
      return <img src={process.env.PUBLIC_URL + '/images/defaultUsuario.png'} alt="" style={this.getImageStyle()} />
    }
    return <img src={this.props.obj.FotoPerfil} alt="" style={this.getImageStyle()} />
  }

  render() {
    return (
      <div>
        <div style={{ padding: '8px 0px 8px 10px' }}>
          <div style={this.ContainerImagem()}>
            {this.getFotoPerfil()}
          </div>
        </div>
      </div>
    )
  }
}

ImagemPerfil.propTypes = {
  obj: PropTypes.object.isRequired,
  ehGrupo: PropTypes.bool,
}

export default ImagemPerfil
