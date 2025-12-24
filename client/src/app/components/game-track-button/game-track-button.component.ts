import { Component, Input } from '@angular/core';
import { HttpHandler } from '../../services/http-handler.service';
import { Game } from '../../interfaces/game';
import { MatButtonModule } from '@angular/material/button';

@Component({
    selector: 'game-track-button',
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
        this.http.postTrackRequest(Number(id)).subscribe({
            next: () => {},
            error: () => {},
        });
    }

    untrack() {
        const id = this.game?.id ?? null;
        if (!id) return;
        this.http.postUntrackRequest(Number(id)).subscribe({
            next: () => {},
            error: () => {},
        });
    }
}
