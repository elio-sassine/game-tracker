export interface Game {
    id: string;
    name: string;
    aggregated_rating: number;
    cover: Cover;
}

export interface Cover {
    id: string;
    url: string;
}
