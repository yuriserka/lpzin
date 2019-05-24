/* eslint-disable require-jsdoc */
import {
  get,
} from 'axios';

export function getUser(userID) {
  get('/usuarios/' + userID)
      .then(
          (result) => {
            console.log(result);
          },
          (error) => {
              console.log(error.message);
          }
      );
}

export function getAllUsers() {
    get('/usuarios')
      .then(
          (result) => {
            console.log(result);
          },
          (error) => {
              console.log(error.message);
          }
      );
}
