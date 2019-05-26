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

export {SearchWrap, SearchBox, SearchTerm};
