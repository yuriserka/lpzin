import styled from 'styled-components';

const SearchWrap = styled.div`
    margin-top: 10px;
    width: 95%;
    position: absolute;
    left: 50%;
    transform: translate(-50%, -50%);
`;

const SearchBox = styled.form`
    width: 100%;
    position: absolute;
    display: flex;
`;

const SearchTerm = styled.input`
    background: whitesmoke;
    width: 100%;
    border: 1px solid gainsboro;
    padding: 5px;
    height: 20px;
    border-radius: 5px;
    outline: none;
    color: black;
`;

// const SearchButton = styled.form`
//     width: 36px;
//     height: 32px;
//     border: 1px solid #00B4CC;
//     background: #921840;
//     color: #fff;
//     border-radius: 0 5px 5px 0;
//     cursor: pointer;
// `;

export {SearchWrap, SearchBox, SearchTerm};
