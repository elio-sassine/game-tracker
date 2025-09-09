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
            .subscribe((usr: User | null) => {
                if (usr == null) {
                    this.errorHandler.showError('An error has occured.');
                } else {
                    this.router.navigate([`/user/${usr.id}`]);
                }
            });
    }
}
