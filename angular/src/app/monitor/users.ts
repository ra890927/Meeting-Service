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
    tag: { name: string, selected: boolean, color: string }[] = [
    ];
    details: string = '';
}