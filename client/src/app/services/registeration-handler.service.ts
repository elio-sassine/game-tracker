import { Injectable } from '@angular/core';
import { ErrorHandler } from './errors.service';
import { HttpHandler } from './http-handler.service';
import { catchError, Observable, Subject } from 'rxjs';
import { HttpErrorResponse, HttpStatusCode } from '@angular/common/http';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Injectable({ providedIn: 'root' })
export class RegistrationHandler {
    constructor(
        private httpHandler: HttpHandler,
        private errorHandler: ErrorHandler,
        private router: Router,
        private snackBar: MatSnackBar
    ) {}

    public makeUser(email: string, password: string, username: string) {
        const response = this.httpHandler.postRegisterRequest(
            email,
            password,
            username
        );
        response
            .pipe(
                catchError((err: HttpErrorResponse) => {
                    if (err.status == HttpStatusCode.Conflict) {
                        this.errorHandler.showError(
                            'An account with this email already exists'
                        );
                    } else {
                        this.errorHandler.showError('Server error');
                    }
                    throw err;
                })
            )
            .subscribe(() => {
                this.router.navigate(['/login']);
            });
    }
}
