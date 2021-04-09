import os
import time

import psycopg2
import requests

API_KEY = os.environ.get('API_KEY')
MOVIE_API_PATH = 'https://kinopoiskapiunofficial.tech/api/v2.1/films'
MOVIE_API_APPEND = '?append_to_response=BUDGET&append_to_response=RATING'
STAFF_API_PATH = 'https://kinopoiskapiunofficial.tech/api/v1/staff'

START_MOVIE_INDEX = 300
END_MOVIE_INDEX = 400


def format_array(arr):
    return '{{{0}}}'.format(', '.join([f'"{x}"' for x in arr]))


def filter_by_profession(people, profession):
    return [person['nameRu'] for person in people if person['professionKey'] == profession]


def get_staff_info(movie_id):
    staff_array = requests.get(f'{STAFF_API_PATH}?filmId={movie_id}',
                               headers={'X-API-KEY': API_KEY}).json()
    return [
        filter_by_profession(staff_array, 'DIRECTOR')[0],
        filter_by_profession(staff_array, 'WRITER')[0],
        filter_by_profession(staff_array, 'PRODUCER')[0],
        filter_by_profession(staff_array, 'OPERATOR')[0],
        filter_by_profession(staff_array, 'COMPOSER')[0],
        filter_by_profession(staff_array, 'DESIGN')[0],
        filter_by_profession(staff_array, 'EDITOR')[0],
        filter_by_profession(staff_array, 'ACTOR'),
    ]


def get_frames_info(movie_id):
    frames_array = requests.get(f'{MOVIE_API_PATH}/{movie_id}/frames',
                                headers={'X-API-KEY': API_KEY}).json()
    frame = frames_array['frames'][0]['image']
    return frame


def get_youtube_thumbnail(youtube_id):
    return f'https://img.youtube.com/vi/{youtube_id}/0.jpg'


def get_trailer_info(movie_id):
    trailers_array = requests.get(f'{MOVIE_API_PATH}/{movie_id}/videos',
                                  headers={'X-API-KEY': API_KEY}).json()
    youtube_trailer_link = [trailer['url'] for trailer in trailers_array['trailers']
                            if trailer['site'] in ['YOUTUBE', 'YouTube']][0]

    if 'v/' in youtube_trailer_link:
        youtube_id = youtube_trailer_link.split('/')[-1]
    else:
        youtube_id = youtube_trailer_link.split('?v=')[1]

    return get_youtube_thumbnail(youtube_id)


def get_movie_info(movie_id):
    response = requests.get(f'{MOVIE_API_PATH}/{movie_id}{MOVIE_API_APPEND}',
                            headers={'X-API-KEY': API_KEY}).json()

    data = response['data']
    rating = response['rating']
    budget = response['budget']

    staff = get_staff_info(movie_id)
    frame = get_frames_info(movie_id)
    trailer_thumbnail = get_trailer_info(movie_id)

    return [
        data['nameRu'],
        data['description'],
        data['year'],
        format_array([item['country'] for item in data['countries']]),
        format_array([item['genre'] for item in data['genres']]),
        data['slogan'],
        *staff[:-1],
        budget['budget'],
        data['filmLength'],
        staff[-1][:5],
        data['posterUrlPreview'],
        frame,
        trailer_thumbnail,
        rating['rating'],
        rating['ratingVoteCount']
    ]


def main():
    conn = psycopg2.connect(
        host='localhost',
        user='mdb',
        password='mdb',
        database='mdb'
    )
    cursor = conn.cursor()

    counter = 0
    for index in range(START_MOVIE_INDEX, END_MOVIE_INDEX):
        try:
            info = [item if item is not None else 'нет данных' for item in get_movie_info(index)]
            cursor.execute(
                'INSERT INTO movie (title, description, productionyear, country, genre, slogan, '
                'director, scriptwriter, producer, operator, composer, artist, montage, '
                'budget, duration, actors, poster, banner, trailerpreview, rating, rating_count) '
                'VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)',
                info)
        except Exception as e:
            print('id:', index, 'error:', e)
            print(e)
            conn.rollback()
            continue

        counter += 1
        conn.commit()
        time.sleep(0.2)

    cursor.execute(
        'INSERT INTO meta (version, movies_count, users_count)'
        'VALUES(%s, %s, %s)', (1, counter, 0))
    conn.commit()

    cursor.close()
    conn.close()


if __name__ == '__main__':
    main()