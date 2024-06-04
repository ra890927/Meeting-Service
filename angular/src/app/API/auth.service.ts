import { Injectable} from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import { environment } from '../../environments/environment';

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
    return this.http.post(environment.apiUrl + 'auth/login', 
    {
      email,
      password
    }, 
    httpOptions);
    
  }
  

  whoami(): Observable<any> {
    return this.http.get(environment.apiUrl + 'auth/whoami', httpOptions);
  }
  
  // logout
  logout(): Observable<any> {
    return this.http.post(environment.apiUrl + 'auth/logout', {}, httpOptions);
  }
}
