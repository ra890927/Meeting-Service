import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})

export class UserService {

  constructor() { }
  
  clean(): void {
    localStorage.removeItem('user');
  }
  // save user data
  public saveUser(user: any): void {
    localStorage.removeItem('user');
    localStorage.setItem('user', JSON.stringify(user));
  }
  // get user data
  public getUser(): any {
    return JSON.parse(localStorage.getItem('user') || '{}');
  }
  //check if user is logged in
  public isLoggedIn(): boolean {
    const user = localStorage.getItem('user');
    if (user) {
      return true;
    }

    return false;
  }
}
