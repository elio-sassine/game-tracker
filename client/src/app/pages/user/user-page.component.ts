import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from '../../interfaces/user';
import { HttpHandler } from '../../services/http-handler.service';

@Component({
    selector: 'user-page',
    templateUrl: 'user-page.component.html',
    styleUrl: 'user-page.component.scss',
    imports: [],
})
export class UserPage implements OnInit {
    userId!: string;
    user!: User;
    constructor(
        private routes: ActivatedRoute,
        private httpHandler: HttpHandler
    ) {}

    ngOnInit(): void {
        this.userId = this.routes.snapshot.paramMap.get('id') as string;
        this.httpHandler
            .getUserRequest(this.userId)
            .subscribe((usr: User | null) => {
                if (usr) {
                    this.user = usr;
                }
            });
    }
}
