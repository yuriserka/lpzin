/* eslint-disable require-jsdoc */
import React from 'react';
import {LogoDiv, MainHDiv, MenuDiv} from './styles.jsx';

class Logo extends React.Component {
  render() {
    return (
      <LogoDiv>
                Menu
      </LogoDiv>
    );
  }
}

class Main extends React.Component {
  render() {
    return (
      <MainHDiv>
        <div>
            Eric visto por ultimo hoje
        </div>
      </MainHDiv>
    );
  }
}

class Header extends React.Component {
  render() {
    return (
      <MenuDiv>
        <Logo />
        <Main />
      </MenuDiv>
    );
  }
}

export default Header;
