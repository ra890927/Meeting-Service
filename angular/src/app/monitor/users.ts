// export class users {
//     id: number = 0;
//     userName: string = '';
//     email: string = '';
//     role: 'admin' | 'user'  = 'user';
// }

export interface rooms {
    id: number;
    roomNumber: string;
    tags: string[];
    capacity: number;
}

export interface users {
    id: number ;
    userName: string ;
    email: string ;
    role: 'admin' | 'user' ;
}

export interface tags {
    id: number;
    tag: string,
    description: string;
}
