//
//  FeedItemCollectionViewCell.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 24.05.2023.
//

import AutoLayoutSugar
import Engine
import UIKit

final class FeedItemCollectionViewCell: UICollectionViewCell, DownloadableimagewidgetDisplayProtocol {
    lazy var widget: DownloadableimagewidgetWidget? = DownloadableimagewidgetCreate(self)

    private let imageView = UIImageView(frame: .zero).prepareForAutoLayout()
    private let captionLabel = UILabel().prepareForAutoLayout()
    private let titleLabel = UILabel().prepareForAutoLayout()

    private var downloadImageTaskID: Int64?

    var viewModel: FeedItemViewModel? {
        didSet {
            updateUI()
        }
    }

    override init(frame: CGRect) {
        super.init(frame: frame)

        setupUI()
        setupLayout()
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    func free() {
        widget?.free()
        widget = nil
    }

    func reset() {
        cancelDownloadImage()
    }

    func downloadImage() {
        cancelDownloadImage()
        downloadImageTaskID = widget?.downloadImageAsync(viewModel?.imageURL)
    }

    // MARK: - Private

    private func setupUI() {
        contentView.addSubview(imageView)
        contentView.addSubview(captionLabel)
        contentView.addSubview(titleLabel)
    }

    private func setupLayout() {
        imageView.topAnchor ~ contentView.topAnchor + 8
        imageView.bottomAnchor ~ contentView.bottomAnchor - 8
        imageView.leadingAnchor ~ contentView.leadingAnchor + 16
        imageView.widthAnchor ~ imageView.heightAnchor

        captionLabel.topAnchor ~ contentView.topAnchor + 8
        captionLabel.leadingAnchor ~ imageView.trailingAnchor + 8

        titleLabel.topAnchor ~ captionLabel.bottomAnchor + 8
        titleLabel.leadingAnchor ~ imageView.trailingAnchor + 8
        titleLabel.trailingAnchor ~ contentView.trailingAnchor - 16
    }

    private func updateUI() {
        guard let viewModel = viewModel else { return }

        captionLabel.text = viewModel.caption
        titleLabel.text = viewModel.title
    }

    private func cancelDownloadImage() {
        guard let downloadImageTaskID = downloadImageTaskID else { return }

        widget?.cancelTask(downloadImageTaskID)
        self.downloadImageTaskID = nil
    }

    // MARK: - DownloadableimagewidgetDisplayProtocol

    func displayImage(_ data: Data?) {
        guard
            let data = data,
            let image = UIImage(data: data)
        else {
            return
        }
        mainAsync { [weak self] in
            self?.imageView.image = image
        }
    }
}
