import { Injectable} from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';


const AUTH_API = 'http://localhost:8080/api/auth/';//need to change

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }
  // login
  // {
  //   "data": {
  //     "message": "string",
  //     "token": "string",
  //     "user": {
  //       "created_at": "string",
  //       "email": "string",
  //       "id": 0,
  //       "role": "string",
  //       "updated_at": "string",
  //       "username": "string"
  //     }
  //   },
  //   "status": "string"
  // }
  login(username: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'login', 
    {
      username,
      password
    }, 
    httpOptions);
  }
  // register
  register(username: string, email: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'register', 
    {
      username,
      email,
      password
    }, 
    httpOptions);
  }

  // logout
  logout(): Observable<any> {
    return this.http.post(AUTH_API + 'logout', {}, httpOptions);
  }
}
