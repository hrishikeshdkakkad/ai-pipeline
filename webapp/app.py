from flask import Flask, render_template, request

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

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
    audio_file.save(f'uploads/{audio_file.filename}')
    video_file.save(f'uploads/{video_file.filename}')

    return render_template('index.html', message='Files uploaded successfully')

if __name__ == '__main__':
    app.run(debug=True)


