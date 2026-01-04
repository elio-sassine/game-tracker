import { Component, OnInit, computed } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from '../../interfaces/user';
import { HttpHandler } from '../../services/http-handler.service';
import { UserGameComponent } from '../../components/user-games/user-game.component';
import { UserInfoComponent } from '../../components/user-info/user-info.component';

@Component({
    selector: 'app-user-page',
    templateUrl: 'user-page.component.html',
    styleUrls: ['user-page.component.scss'],
    imports: [UserGameComponent, UserInfoComponent],
})
export class UserPageComponent implements OnInit {
    userId!: string;
    user!: User;
    loading = true;
    loadingState = computed(() => this.loading);

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
