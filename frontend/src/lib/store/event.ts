import type { EventBase } from '$lib/model/EventBase';
import { writable } from 'svelte/store';

export const event = writable<EventBase>();
