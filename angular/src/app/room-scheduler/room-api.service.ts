import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class RoomAPIService {
  constructor(private http: HttpClient) { }
  getRooms() {
    return this.http.get('https://jsonplaceholder.typicode.com/todos');
  }
  getRoom(id: number) {
    return this.http.get('http://localhost:3000/rooms/' + id);
  }

  getRoomEvents(id: number) {
    return this.http.get('http://localhost:3000/rooms/' + id + '/events');
  }

  addEvent(id: number, event: any) {
    return this.http.post('http://localhost:3000/rooms/' + id + '/events', event);
  }

  updateEvent(id: number, event: any) {
    return this.http.put('http://localhost:3000/rooms/' + id + '/events/' + event.id, event);
  }

  deleteEvent(id: number, eventId: number) {
    return this.http.delete('http://localhost:3000/rooms/' + id + '/events/' + eventId);
  }
}