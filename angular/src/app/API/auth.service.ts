import { Injectable} from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';

const AUTH_API = 'http://140.113.215.132:8080/api/v1/auth/';

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
  login(email: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'login', 
    {
      email,
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

  whoami(): Observable<any> {
    return this.http.get(AUTH_API + 'whoami', httpOptions);
  }
  
  // logout
  logout(): Observable<any> {
    return this.http.post(AUTH_API + 'logout', {}, httpOptions);
  }
}
