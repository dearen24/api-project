import con from "../database/database.js";

export async function getUser(username) {
  return new Promise((resolve, reject) => {
    con.query(`SELECT * FROM users WHERE username = '${username}'`, (err, result) => {
      if(err){
        reject(err);
      }
      else{
        resolve(JSON.parse(JSON.stringify(result)));
      }
    });
  })
}

export async function getOneUser(id) {
  return new Promise((resolve, reject) => {
    con.query(`SELECT * FROM users WHERE id = '${id}'`, (err, result) => {
      if(err){
        reject(err);
      }
      else{
        resolve(JSON.parse(JSON.stringify(result)));
      }
    });
  })
}

export async function getUsers(){
  return new Promise((resolve, reject) => {
    con.query(`SELECT * FROM users`, (err, result) => {
      if(err){
        reject(err);
      }
      else{
        resolve(JSON.parse(JSON.stringify(result)));
      }
    });
  })
}

export async function createUser(username, email, password, firstName, lastName, isActive){
  return new Promise((resolve, reject) => {
    con.query(`INSERT INTO users (username, email, password_hash, first_name, last_name, is_active, created_at, updated_at) VALUES('${username}','${email}','${password}','${firstName}','${lastName}',${isActive}, NOW(), NOW())`, (err, result) => {
      if(err){
        reject(err);
      }
      else{
        resolve(JSON.parse(JSON.stringify(result)));
      }
    });
  })
}

export async function editUser(id, username, email, password, firstName, lastName, isActive){
  return new Promise((resolve, reject) => {
    con.query(`UPDATE users SET username = '${username}', email = '${email}', password_hash = '${password}', first_name = '${firstName}', last_name = '${lastName}', is_active = ${isActive}, updated_at = NOW() WHERE id=${id}`, (err, result) => {
      if(err){
        reject(err);
      }
      else{
        resolve(JSON.parse(JSON.stringify(result)));
      }
    });
  })
}

export async function deleteUser(id){
  return new Promise((resolve, reject) => {
    con.query(`DELETE FROM users WHERE id=${id}`, (err, result) => {
      if(err){
        reject(err);
      }
      else{
        resolve(JSON.parse(JSON.stringify(result)));
      }
    });
  })
}