import styled from 'styled-components';

const MenuDiv = styled.div`
  color: white;
  background: #921840;
  height: 48px;
  margin: 0;
`;

const LogoDiv = styled(MenuDiv)`
  float: left;
  width: 31%;
  &:hover {
    background: #5a0e27;
  }
`;

const MainHDiv = styled(MenuDiv)`
  float:left;
  width: 69%;
  &:hover {
    background: #5a0e27;
  }
`;

export {MenuDiv, LogoDiv, MainHDiv};
