import os
from flask import Flask, render_template, request, redirect, url_for, send_file,jsonify
from subprocess import Popen,call
import requests

app = Flask(__name__)


@app.route('/')
def index():
    # Get a list of the filenames in the "uploads" folder
    filenames = os.listdir('uploads/video')
    return render_template('index.html', filenames=filenames)

@app.route('/upload', methods=['POST'])
def upload():
    audio_file = request.files['audioFile']
    video_file = request.files['videoFile']

    if not audio_file or not video_file:
        return render_template('index.html', message='Please select both audio and video files')

    allowed_audio_formats = {'mp3', 'wav'}
    allowed_video_formats = {'mp4', 'avi'}

    audio_extension = audio_file.filename.rsplit('.', 1)[-1].lower()
    video_extension = video_file.filename.rsplit('.', 1)[-1].lower()

    if audio_extension not in allowed_audio_formats or video_extension not in allowed_video_formats:
        return render_template('index.html', message='Please upload files with allowed formats: mp3, wav, mp4, avi')

    # save files to separate folders
    audio_file.save(f'uploads/audio/{audio_file.filename}')
    video_file.save(f'uploads/video/{video_file.filename}')

    return redirect(url_for('list_videos'))

@app.route('/list_videos')
def list_videos():
    # Get a list of the filenames in the "uploads/video" folder
    filenames = os.listdir('uploads/video/')
    return render_template('list_videos.html', filenames=filenames)

@app.route('/play/<filename>')
def play(filename):
    # Check if the file exists
    print(filename)
    if not os.path.isfile('uploads/video/' + filename):
        return "File not found"

    # Play the video using ffplay
    call(['ffplay', '-i', 'uploads/video/' + filename])

    return ''

@app.route('/clanText')
def get_clan_text():
    response = requests.get('http://localhost:3028/clanText')
    data = response.content
    return render_template('clan_text.html', data=data.decode())

if __name__ == '__main__':
    app.run(debug=True)
