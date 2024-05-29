import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
interface User {
  id: number;
  username: string;
  email: string;
}
@Injectable({
  providedIn: 'root'
})

export class UserService {

  constructor() { }
  
  clean(): void {
    window.sessionStorage.removeItem('user');
    window.sessionStorage.removeItem('token');
  }
  // save user data
  public saveUser(user: any): void {
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
