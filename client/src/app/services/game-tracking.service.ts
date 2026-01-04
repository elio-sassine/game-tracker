import { Injectable } from '@angular/core';
import { HttpHandler } from './http-handler.service';
import { User } from '../interfaces/user';

@Injectable({
    providedIn: 'root',
})
export class GameTrackingService {
    constructor(private http: HttpHandler) {}

    public getUserTrackedGames(user: User) {
        return this.http.getTrackedGamesRequest(user.id ?? '');
    }
}
