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
    tags: string[]=[];
    capacity: number = 0;
    details: string = '';
}

export const allTags: string[] = ['Projector Available', 'Free WiFi', 'Air Conditioning', 'Food Allowed', 'Whiteboard'];
