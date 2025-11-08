import { Router } from '@angular/router';
import { ErrorHandler } from './errors.service';
import { HttpHandler } from './http-handler.service';
import { Injectable } from '@angular/core';
import { catchError } from 'rxjs';
import { HttpErrorResponse, HttpStatusCode } from '@angular/common/http';
import { User } from '../interfaces/user';

@Injectable({
    providedIn: 'root',
})
export class LoginHandler {
    constructor(
        private httpHandler: HttpHandler,
        private errorHandler: ErrorHandler,
        private router: Router
    ) {}

    public loginUser(email: string, password: string) {
        this.httpHandler
            .postLoginRequest(email, password)
            .pipe(
                catchError((err: HttpErrorResponse) => {
                    if (err.status == HttpStatusCode.NotFound) {
                        this.errorHandler.showError(
                            'This email or password is not valid.'
                        );
                    }
                    throw err;
                })
            )
            .subscribe((id: string | null) => {
                if (id == null) {
                    this.errorHandler.showError('No ID!');
                } else {
                    console.log('User logged in: ', id);
                    this.router.navigate([`/user/${id}`]);
                }
            });
    }
}
