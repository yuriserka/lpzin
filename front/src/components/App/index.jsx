/* eslint-disable require-jsdoc */
import React from 'react';
import Header from '../Header/index';
import SideBar from '../Sidebar/index';
import MsgBox from '../MsgBox/index';
import PageStyle from './styles.jsx';

class App extends React.Component {
  render() {
    return (
      <PageStyle>
        <Header />
        <MsgBox />
        <SideBar />
      </PageStyle>
    );
  }
}

export default App;
