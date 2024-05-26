export class users {
    id: string = '';
    userName: string = '';
    email: string = '';
    role: 'Admin' | 'User'  = 'User';
    details: string = '';
}

export class rooms {
    id: string = '';
    roomNumber: string = '';
    fruits: string[]=[];
    capacity: number = 0;
    details: string = '';
}