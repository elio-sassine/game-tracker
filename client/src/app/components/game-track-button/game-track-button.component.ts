import { Component, Input } from '@angular/core';
import { HttpHandler } from '../../services/http-handler.service';
import { Game } from '../../interfaces/game';
import { MatButtonModule } from '@angular/material/button';

@Component({
    selector: 'app-game-track-button',
    standalone: true,
    imports: [MatButtonModule],
    templateUrl: './game-track-button.component.html',
    styleUrls: ['./game-track-button.component.scss'],
})
export class GameTrackButtonComponent {
    @Input() game?: Game | null;

    constructor(private http: HttpHandler) {}

    track() {
        const id = this.game?.id ?? null;
        if (!id) return;
        const idString = id.toString();
        this.http.postTrackRequest(idString).subscribe({
            next: () => {
                return;
            },
            error: () => {
                return;
            },
        });
    }

    untrack() {
        const id = this.game?.id ?? null;
        if (!id) return;
        const idString = id.toString();
        this.http.postUntrackRequest(idString).subscribe({
            next: () => {
                return;
            },
            error: () => {
                return;
            },
        });
    }
}
