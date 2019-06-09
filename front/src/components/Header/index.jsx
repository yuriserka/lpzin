/* eslint-disable require-jsdoc */
import React from 'react';
import {MenuDiv} from './styles.jsx';
import Logo from './Logo/index';

class Header extends React.Component {
  render() {
    return (
      <MenuDiv>
        <Logo />
      </MenuDiv>
    );
  }
}

export default Header;
