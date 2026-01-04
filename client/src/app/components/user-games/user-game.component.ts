import { Component, computed, input } from '@angular/core';
import { User } from '../../interfaces/user';
import { GameTrackingService } from '../../services/game-tracking.service';
import { Game } from '../../interfaces/game';
import { GameComponent } from '../game/game.component';
import { effect, signal } from '@angular/core';

@Component({
    selector: 'app-user-game',
    templateUrl: 'user-game.component.html',
    styleUrl: 'user-game.component.scss',
    imports: [GameComponent],
})
export class UserGameComponent {
    userInput = input<User>();

    user = computed(() => this.userInput());

    constructor(private gameTrackingService: GameTrackingService) {}

    // signal that holds the fetched games
    trackedGames = signal<Game[]>([]);

    // fetch when user changes
    private fetch = effect(() => {
        const u = this.user();
        if (!u) {
            this.trackedGames.set([]);
            return;
        }
        this.gameTrackingService.getUserTrackedGames(u).subscribe({
            next: (res) => {
                this.trackedGames.set(res);
            },
            error: () => this.trackedGames.set([]),
        });
    });
}
