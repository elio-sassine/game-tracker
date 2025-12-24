import { HttpClient } from '@angular/common/http';
import { Injectable, Inject } from '@angular/core';
import { Observable } from 'rxjs';
import { Game } from '../interfaces/game';
import { User } from '../interfaces/user';

@Injectable({
    providedIn: 'root',
})
export class HttpHandler {
    constructor(@Inject(HttpClient) private http: HttpClient) {}

    baseUrl = 'https://localhost:8443';
    jwtKey = null;

    getGamesRequest(name: string): Observable<Game[]> {
        return this.http.get<Game[]>(`${this.baseUrl}/search?name=${name}`);
    }

    postRegisterRequest(
        email: string,
        password: string,
        username: string
    ): Observable<unknown> {
        return this.http.post<unknown>(`${this.baseUrl}/register`, {
            email: email,
            username: username,
            password: password,
        });
    }

    postLoginRequest(
        email: string,
        password: string
    ): Observable<string | null> {
        // Return the user ID as a string to preserve full 64-bit snowflake IDs
        return this.http.post<string | null>(
            `${this.baseUrl}/login`,
            {
                email: email,
                password: password,
            },
            { withCredentials: true }
        );
    }

    getUserRequest(id: string): Observable<User | null> {
        // id passed as string to avoid precision loss in JavaScript numbers
        return this.http.get<User | null>(`${this.baseUrl}/user?id=${id}`, {
            withCredentials: true,
        });
    }

    postTrackRequest(gameId: number): Observable<unknown> {
        return this.http.post(
            `${this.baseUrl}/gameTracking/track`,
            { game: gameId.toString() },
            { withCredentials: true }
        );
    }

    postUntrackRequest(gameId: number): Observable<unknown> {
        return this.http.post(
            `${this.baseUrl}/gameTracking/untrack`,
            { game: gameId.toString() },
            { withCredentials: true }
        );
    }
}
