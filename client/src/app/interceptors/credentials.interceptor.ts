import { Injectable } from '@angular/core';
import {
    HttpEvent,
    HttpHandler,
    HttpInterceptor,
    HttpRequest,
} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class CredentialsInterceptor implements HttpInterceptor {
    intercept(
        req: HttpRequest<unknown>,
        next: HttpHandler
    ): Observable<HttpEvent<unknown>> {
        const cloned = req.clone({ withCredentials: true });
        return next.handle(cloned);
    }
}
