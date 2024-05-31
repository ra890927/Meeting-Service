// export class users {
//     id: number = 0;
//     userName: string = '';
//     email: string = '';
//     role: 'admin' | 'user'  = 'user';
// }

export class rooms {
    id: number = 0;
    roomNumber: string = '';
    tags: string[]=[];
    capacity: number = 0;
    // details: string = '';
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

export const allTags: tags[] = [
    { id: 0, tag: 'No Smoking', description: '禁止吸菸'},
    { id: 1, tag: 'Food Allowed', description: ''},
    { id: 2, tag: 'Projector Available', description: ''},
    { id: 3, tag: 'Air Conditioning', description: ''},

    // 'Projector Available', 'Free WiFi', 'Air Conditioning', 'Food Allowed', 'Whiteboard'
];
