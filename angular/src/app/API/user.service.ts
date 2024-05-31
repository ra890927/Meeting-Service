import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

const AUTH_API = 'http://140.113.215.132:8080/api/v1/user';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

interface User {
  id: number;
  username: string;
  email: string;
}
@Injectable({
  providedIn: 'root'
})



export class UserService {

  constructor( private http: HttpClient ) {}

  // register
  register(username: string, email: string, password: string): Observable<any> {
    return this.http.post(AUTH_API, 
    {
      username,
      email,
      password
    }, 
    httpOptions);
  }
  
  clean(): void {
    window.sessionStorage.removeItem('user');
    window.sessionStorage.removeItem('token');
  }
  // save user data
  public saveUser(user: any): void {
    console.log(user);
    window.sessionStorage.removeItem('user');
    window.sessionStorage.setItem('user', JSON.stringify(user));
  }
  // get user data
  public getUser(): User|null{
    const user = window.sessionStorage.getItem('user');
    if (user) {
      return JSON.parse(user);
    }
    return null;
  }
  //check if user is logged in
  public isLoggedIn(): boolean {
    const user = window.sessionStorage.getItem('user');
    if (user) {
      return true;
    }

    return false;
  }
}
