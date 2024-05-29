// export class users {
//     id: number = 0;
//     userName: string = '';
//     email: string = '';
//     role: 'admin' | 'user'  = 'user';
// }

export class rooms {
    id: string = '';
    roomNumber: string = '';
    tags: string[]=[];
    capacity: number = 0;
    details: string = '';
}

export interface users {
    id: number ;
    userName: string ;
    email: string ;
    role: 'admin' | 'user' ;
}

export const allTags: string[] = ['Projector Available', 'Free WiFi', 'Air Conditioning', 'Food Allowed', 'Whiteboard'];
